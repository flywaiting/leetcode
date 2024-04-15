struct MyHashSet {
    list: Vec<bool>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyHashSet {
    fn new() -> Self {
        MyHashSet {
            list: vec![false; 1000_001],
        }
    }

    fn add(&mut self, key: i32) {
        self.list[key as usize] = true;
    }

    fn remove(&mut self, key: i32) {
        self.list[key as usize] = false;
    }

    fn contains(&self, key: i32) -> bool {
        self.list[key as usize]
    }
}
