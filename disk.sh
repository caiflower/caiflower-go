#!/bin/bash

#set -x

readonly IGNORE_PART="-1 NULL NULL"

# 检查并安装依赖组件
check_and_install() {
    local package=$1
    local command=$2
    if ! command -v $command &> /dev/null; then
        echo "Installing $package..."
        if command -v apt-get &> /dev/null; then
            apt-get update
            apt-get install -y $package
        elif command -v yum &> /dev/null; then
            yum install -y $package
        elif command -v dnf &> /dev/null; then
            dnf install -y $package
        else
            echo "Error: Package manager not found. Please install $package manually."
            exit 1
        fi
    fi
}

# 检查并安装所需的依赖组件
check_and_install parted parted
check_and_install util-linux lsblk
check_and_install util-linux blkid
check_and_install util-linux wipefs
check_and_install coreutils numfmt
check_and_install e2fsprogs mkfs.ext4
check_and_install xfsprogs mkfs.xfs

# 定义动作
ACTION="{{ACTION}}"
# 定义磁盘设备
DISK="{{DISK}}"

# 定义每个分区的信息
PART1="{{PART1}}"
PART2="{{PART2}}"
PART3="{{PART3}}"
PART4="{{PART4}}"
PART5="{{PART5}}"

# 检查每个分区信息变量是否非空，并将其添加到数组中

if [[ -n "$PART1" && "$PART1" =~ [^[:space:]] && "$PART1" != "$IGNORE_PART" ]]; then
    PARTITIONS+=("$PART1")
fi

if [[ -n "$PART2" && "$PART2" =~ [^[:space:]] && "$PART2" != "$IGNORE_PART" ]]; then
    PARTITIONS+=("$PART2")
fi

if [[ -n "$PART3" && "$PART3" =~ [^[:space:]] && "$PART3" != "$IGNORE_PART" ]]; then
    PARTITIONS+=("$PART3")
fi

if [[ -n "$PART4" && "$PART4" =~ [^[:space:]] && "$PART4" != "$IGNORE_PART" ]]; then
    PARTITIONS+=("$PART4")
fi

if [[ -n "$PART5" && "$PART5" =~ [^[:space:]] && "$PART5" != "$IGNORE_PART" ]]; then
    PARTITIONS+=("$PART5")
fi

# 检查磁盘是否存在
if [ ! -b "$DISK" ]; then
    echo "Error: Disk $DISK does not exist."
    exit 1
fi

if [ "$ACTION" == "init" ]; then
    # 获取磁盘信息
    DISK_SIZE=$(lsblk -bno SIZE $DISK | head -n 1)
    if [ $? -ne 0 ]; then
        echo "Error: Failed to get disk size for $DISK"
        exit 1
    fi

    SECTOR_SIZE=$(cat /sys/block/$(basename $DISK)/queue/hw_sector_size)
    if [ $? -ne 0 ]; then
        echo "Error: Failed to get sector size for $DISK"
        exit 1
    fi

    TOTAL_SECTORS=$((DISK_SIZE / SECTOR_SIZE))

    # 计算所有分区的总大小
    TOTAL_PARTITION_SIZE=0
    for PARTITION_INFO in "${PARTITIONS[@]}"; do
        SIZE=$(echo $PARTITION_INFO | awk '{print $1}')
        SIZE_IN_BYTES=$(numfmt --from=iec $SIZE)
        TOTAL_PARTITION_SIZE=$((TOTAL_PARTITION_SIZE + SIZE_IN_BYTES))
    done

    # 检查分区总大小是否超过磁盘总空间
    if [ $TOTAL_PARTITION_SIZE -gt $DISK_SIZE ]; then
        echo "Error: Total partition size ($TOTAL_PARTITION_SIZE bytes) exceeds disk size ($DISK_SIZE bytes)."
        exit 1
    fi

    cp /etc/fstab /etc/fstab.init.bak
    parted /dev/vdb --script rm 1

    # 卸载磁盘上的所有分区
    for PART in $(lsblk -ln -o NAME $DISK | grep -E "^${DISK#/dev/}[0-9]+$"); do
        PART_PATH="/dev/$PART"
        MOUNTPOINT=$(findmnt -nr -o TARGET $PART_PATH)
        if [ -n "$MOUNTPOINT" ]; then
            umount $PART_PATH
            if [ $? -ne 0 ]; then
                echo "Error: Failed to unmount $PART_PATH"
                exit 1
            fi
            # 获取分区的 UUID
            UUID=$(blkid -s UUID -o value $PART_PATH)
            if [ -n "$UUID" ]; then
                # 清理 /etc/fstab 中的相关条目
                sed -i "/UUID=$UUID/d" /etc/fstab
                if [ $? -ne 0 ]; then
                    echo "Error: Failed to remove UUID=$UUID from /etc/fstab"
                    exit 1
                fi
            else
                echo "Error: Failed to get UUID for $PART_PATH"
                exit 1
            fi
        fi
        PART_NUM=${PART#${DISK#/dev/}}
        parted $DISK --script rm $PART_NUM
        if [ $? -ne 0 ]; then
            echo "Error: Failed to delete partition $PART"
            exit 1
        fi
    done
    #
    wipefs -a $DISK

    # 创建分区表
    #parted $DISK mklabel gpt
    parted $DISK --script -- mklabel gpt
    if [ $? -ne 0 ]; then
        echo "Error: Failed to create partition table on $DISK"
        exit 1
    fi
    if [ $? -ne 0 ]; then
        echo "Error: Failed to create partition table on $DISK"
        exit 1
    fi

    # 起始扇区
    START_SECTOR=2048

    # 处理每个分区参数
    #for PARTITION_INFO in "${PARTITIONS[@]}"; do
    for i in "${!PARTITIONS[@]}"; do
        PARTITION_INFO=${PARTITIONS[$i]}
        SIZE=$(echo $PARTITION_INFO | awk '{print $1}')
        FILESYSTEM=$(echo $PARTITION_INFO | awk '{print $2}')
        MOUNTPOINT=$(echo $PARTITION_INFO | awk '{print $3}')

        # 计算结束扇区
        SIZE_IN_BYTES=$(numfmt --from=iec $SIZE)
        SECTOR_COUNT=$((SIZE_IN_BYTES / SECTOR_SIZE))
        END_SECTOR=$((START_SECTOR + SECTOR_COUNT - 1))

        # 如果是最后一个分区，调整结束扇区以确保不超出磁盘大小
        if [ $i -eq $((${#PARTITIONS[@]} - 1)) && [ $END_SECTOR -gt $((TOTAL_SECTORS - 33 - 1)) ]; then
            END_SECTOR=$((TOTAL_SECTORS - 33 - 1))
        fi

        # 检查是否超出磁盘大小
        if [ $END_SECTOR -ge $TOTAL_SECTORS ]; then
            echo "Error: The location ${END_SECTOR}s is outside of the device $DISK."
            exit 1
        fi

        # 创建分区
        parted $DISK --script mkpart primary ${START_SECTOR}s ${END_SECTOR}s
        if [ $? -ne 0 ]; then
            echo "Error: Failed to create partition on $DISK"
            exit 1
        fi

        # 等待分区创建完成
        sleep 2s

        # 获取最后一个分区名称
        PARTITION=$(lsblk -no KNAME $DISK | sort -V | tail -n 1)
        if [ -z "$PARTITION" ]; then
            echo "Error: Failed to get partition name for $DISK"
            exit 1
        fi
        PARTITION="/dev/$PARTITION"

        # 清除旧的文件系统信息（可选）
        wipefs --all $PARTITION

        # 格式化分区
        mkfs.$FILESYSTEM $PARTITION
        if [ $? -ne 0 ]; then
            echo "Error: Failed to format partition $PARTITION"
            exit 1
        fi

        # 创建挂载点并挂载分区
        mkdir -p $MOUNTPOINT
        if [ $? -ne 0 ]; then
            echo "Error: Failed to create mount point $MOUNTPOINT"
            exit 1
        fi

        mount $PARTITION $MOUNTPOINT
        if [ $? -ne 0 ]; then
            echo "Error: Failed to mount partition $PARTITION to $MOUNTPOINT"
            exit 1
        fi

        # 将分区信息写入/etc/fstab
        UUID=$(blkid -s UUID -o value $PARTITION)
        if [ $? -ne 0 ]; then
            echo "Error: Failed to get UUID for partition $PARTITION"
            exit 1
        fi
        # 删除已存在的相同UUID数据
        sed -i "/UUID=$UUID/d" /etc/fstab

        echo "UUID=$UUID $MOUNTPOINT $FILESYSTEM defaults 0 0" >> /etc/fstab
        if [ $? -ne 0 ]; then
            echo "Error: Failed to write to /etc/fstab"
            exit 1
        fi

        # 更新起始扇区
        START_SECTOR=$((END_SECTOR + 1))
        # 确保起始扇区对齐
        parted $DISK --script align-check optimal $((i + 1))
        if [ $? -ne 0 ]; then
            echo "Error: Partition $((i + 1)) is not aligned properly."
            exit 1
        fi
    done

    # 禁用调试模式
    set +x

    echo "Disk partitioning and setup completed successfully."

elif [ "$ACTION" == "mount" ]; then
    # 获取现有分区
    EXISTING_PARTITIONS=$(lsblk -ln -o NAME $DISK | grep -E "^${DISK#/dev/}[0-9]+$")

    # 检查是否有足够的现有分区
    if [ $(echo "$EXISTING_PARTITIONS" | wc -l) -lt ${#PARTITIONS[@]} ]; then
        echo "Error: Not enough existing partitions on $DISK"
        exit 1
    fi

    # 处理每个分区信息
    for i in "${!PARTITIONS[@]}"; do
        PARTITION_INFO=${PARTITIONS[$i]}
        FILESYSTEM=$(echo $PARTITION_INFO | awk '{print $2}')
        MOUNTPOINT=$(echo $PARTITION_INFO | awk '{print $3}')

        # 获取现有分区名称
        PARTITION=$(echo "$EXISTING_PARTITIONS" | sed -n "$((i + 1))p")
        PARTITION="/dev/$PARTITION"

        # 创建挂载点并挂载分区
        mkdir -p $MOUNTPOINT
        if [ $? -ne 0 ]; then
            echo "Error: Failed to create mount point $MOUNTPOINT"
            exit 1
        fi

        mount $PARTITION $MOUNTPOINT
        if [ $? -ne 0 ]; then
            echo "Error: Failed to mount partition $PARTITION to $MOUNTPOINT"
            exit 1
        fi

        # 获取分区的UUID
        UUID=$(blkid -s UUID -o value $PARTITION)
        if [ $? -ne 0 ]; then
            echo "Error: Failed to get UUID for partition $PARTITION"
            exit 1
        fi

        # 删除已存在的相同UUID数据
        sed -i "|UUID=$UUID|d" /etc/fstab

        # 将分区信息写入/etc/fstab
        echo "UUID=$UUID $MOUNTPOINT $FILESYSTEM defaults 0 0" >> /etc/fstab
        if [ $? -ne 0 ]; then
            echo "Error: Failed to write to /etc/fstab"
            exit 1
        fi
    done

    echo "Disk mounting completed successfully."

else
    echo "Invalid action: $ACTION"
    echo "Usage: $0 {init|mount}"
    exit 1
fi