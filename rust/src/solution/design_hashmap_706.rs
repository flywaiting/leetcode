struct MyHashMap {
    list: [i32; 1_000_001],
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyHashMap {
    fn new() -> Self {
        MyHashMap {
            list: [-1; 1_000_001],
        }
    }

    fn put(&mut self, key: i32, value: i32) {
        self.list[key as usize] = value;
    }

    fn get(&self, key: i32) -> i32 {
        self.list[key as usize]
    }

    fn remove(&mut self, key: i32) {
        self.list[key as usize] = -1;
    }
}
