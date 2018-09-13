#[no_mangle]
pub extern "C" fn add_one(x: i32) -> i32 {
    add_another_one(x+1)
}

#[no_mangle]
pub extern "C" fn add_another_one(x: i32) -> i32 {
    x + 1
}
