use super::config::*;

pub struct Solution;

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn sorted_list_to_bst(head: Option<Box<ListNode>>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut len = 0;
        let mut p = &head;
        while let Some(node) = p {
            len += 1;
            p = &node.next;
        }

        let (root, _) = Self::build_tree(head, 0, len - 1);
        root
    }

    fn build_tree(
        head: Option<Box<ListNode>>,
        left: i32,
        right: i32,
    ) -> (Option<Rc<RefCell<TreeNode>>>, Option<Box<ListNode>>) {
        if left > right {
            return (None, head);
        }

        let mid = (left + right + 1) / 2;
        let (ln, head) = Self::build_tree(head, left, mid - 1);
        if let Some(node) = head {
            // head = node.next;
            let mut root = TreeNode::new(node.val);
            root.left = ln;
            let (rn, head) = Self::build_tree(node.next, mid + 1, right);
            root.right = rn;
            return (Some(Rc::new(RefCell::new(root))), head);
        }
        (None, head)
    }
}
