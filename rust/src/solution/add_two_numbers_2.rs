use super::config::*;

// 被这个链表恶心坏了
// option 这个要再琢磨琢磨
// box 这个倒是挺好理解的
// 所有权这块 完全被绕晕了。。。
impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        
        Self::add_numbers(l1,l2, 0)
    }

    pub fn add_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>, mut other:i32) -> Option<Box<ListNode>> {
        if l1.is_none() && l2.is_none() && other ==0 {
            None
        }else {
            Some(Box::new(ListNode { 
                next: Self::add_numbers(
                    l1.and_then(|a| {other+=a.val; a.next}), 
                    l2.and_then(|a| {other+=a.val; a.next}),
                    other/10
                ),
                val: other%10  
            }))
        }
    }
}

pub struct Solution;
