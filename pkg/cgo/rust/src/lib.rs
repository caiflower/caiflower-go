#[no_mangle]
pub extern "C" fn say_hello(name:* const i8,age:std::ffi::c_int){
let rname=unsafe{
let realname=std::ffi::CStr::from_ptr(name);
realname.to_str().expect("to string failed")
};
println!("hello {rname} your age is {age}");
}