use super::config::ListNode;

impl Solution {
    pub fn merge_in_between(
        list1: Option<Box<ListNode>>,
        a: i32,
        b: i32,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut head = Box::new(ListNode {
            next: list1,
            val: 0,
        });
        let mut cur = head.as_mut();

        for _ in 0..a {
            cur = cur.next.as_mut().unwrap();
        }
        let mut p = cur.next.take();
        for _ in a..b {
            p = p.unwrap().next;
        }

        cur.next = list2;
        while cur.next.is_some() {
            cur = cur.next.as_mut().unwrap();
        }
        cur.next = p.unwrap().next.take();

        head.next
    }
}

pub struct Solution;
