use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        let mut lists = lists;
        if lists.is_empty() {
            return None;
        }

        let mut size = lists.len();
        while size > 1 {
            for i in 0..size / 2 {
                lists[i] = Self::merge(lists[i].take(), lists[size - i - 1].take());
            }
            size = (size + 1) / 2;
        }
        lists[0].take()
    }

    fn merge(
        mut l1: Option<Box<ListNode>>,
        mut l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut dummy = Box::new(ListNode::new(0));
        let mut p = &mut dummy;

        while l1.is_some() || l2.is_some() {
            match (l1, l2) {
                (Some(mut node1), Some(mut node2)) => {
                    if node1.val < node2.val {
                        l1 = node1.next.take();
                        l2 = Some(node2);
                        p.next = Some(node1);
                    } else {
                        l1 = Some(node1);
                        l2 = node2.next.take();
                        p.next = Some(node2);
                    }
                    p = p.next.as_mut().unwrap();
                }
                (Some(node), None) | (None, Some(node)) => {
                    l1 = None;
                    l2 = None;
                    p.next = Some(node);
                }
                _ => break,
            }
        }
        dummy.next
    }
}
