package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Student struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Score struct {
			Math    int `json:"math"`
			English int `json:"english"`
			Chinese int `json:"chinese"`
		} `json:"score"`
	} `json:"spec"`
}

type StudentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Student `json:"items"`
}

func (s *Student) DeepCopyObject() runtime.Object {
	if s == nil {
		return nil
	}
	stu := new(Student)
	*stu = *s
	stu.TypeMeta = s.TypeMeta
	s.ObjectMeta.DeepCopyInto(&stu.ObjectMeta)
	return stu
}

func (s *Student) DeepCopyInto(target *Student) {
	*target = *s
	s.ObjectMeta.DeepCopyInto(&target.ObjectMeta)
}

func (sl *StudentList) DeepCopyObject() runtime.Object {
	if sl == nil {
		return nil
	}
	ret := new(StudentList)
	*ret = *sl
	ret.Items = make([]Student, len(sl.Items))
	for i, v := range sl.Items {
		v.DeepCopyInto(&ret.Items[i])
	}
	return ret
}

func init() {
	SchemeBuilder.Register(&Student{}, &StudentList{})
}
