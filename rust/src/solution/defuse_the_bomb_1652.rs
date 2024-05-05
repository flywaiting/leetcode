impl Solution {
    pub fn decrypt(code: Vec<i32>, k: i32) -> Vec<i32> {
        let size = code.len();
        let mut res = vec![0; size];
        if k == 0 {
            return res;
        }

        let m = size as i32;
        let mut i: i32 = 1.min(k);
        while i < 0.max(k) {
            (0..size).for_each(|j| {
                let x = i + (j as i32) + m;
                res[j] += code[(x as usize) % size];
            });
            i += 1;
        }
        res
    }
}

pub struct Solution;
