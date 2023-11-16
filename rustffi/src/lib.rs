use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn basic_string_pass(value: *const libc::c_char) -> *mut libc::c_char {
    let c_str = unsafe { CStr::from_ptr(value) };
    let r_str = c_str.to_str().unwrap();
    let r_str = format!("Hello, {}!", r_str);
    let r_str = std::ffi::CString::new(r_str).unwrap();
    r_str.into_raw()
}

#[no_mangle]
pub extern "C" fn long_running(length: libc::c_int) {
    for i in 0..length {
        println!("{} of {} (rs)", i, length);
        std::thread::sleep(std::time::Duration::from_secs(1));
    }
}
