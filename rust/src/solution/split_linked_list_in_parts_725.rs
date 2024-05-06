use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn split_list_to_parts(head: Option<Box<ListNode>>, k: i32) -> Vec<Option<Box<ListNode>>> {
        let mut head = head;
        let mut list = vec![None; k as usize];
        let mut cnt = 0;
        let mut p = head.as_ref();
        while let Some(node) = p {
            cnt += 1;
            p = node.next.as_ref();
        }

        let mut remain = cnt % k;

        for i in 0..k as usize {
            let mut cnt = cnt / k;
            list[i] = head;

            let mut p = list[i].as_mut();
            if remain > 0 {
                cnt += 1;
                remain -= 1;
            }
            // 计数从 1 开始
            for _ in 1..cnt {
                p = p.unwrap().next.as_mut();
                // p = p.and_then(|node| node.next.as_mut());    // good then before
            }

            if let Some(node) = p {
                head = node.next.take();
            } else {
                break;
            }
        }
        list
    }
}

// 简单的解题思路
// 所有权问题 不熟练 还想着各种引用各司其职。。。
