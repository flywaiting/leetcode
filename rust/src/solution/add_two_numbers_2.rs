use super::config::*;

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let head = Box::new(ListNode::new(0));
        // let tail:Option<Box<ListNode>>;
        // let tail = head.next;
        let mut carry = 0;
        let mut is_end=false;

        let mut t = (l1, l2);
        loop {
            match t {
                (Some(a), Some(b))=>{
                    carry += a.val+b.val;
                    t = (a.next, b.next);
                },
                (Some(a), None)|(None, Some(a))=>{
                    carry += a.val;
                    t = (a.next, None);
                },
                _=> is_end=true;
            }

            if !is_end || carry!=0 {
                let node = Box::new(ListNode::new(carry%10));
                *tail = Some(Box::new(ListNode::new(carry/10)));
                tail = &mut tail.as_mut();

                carry/=10;
            }
            if is_end {
                break;
            }
        };
        head.next
    }
}

pub struct Solution;
