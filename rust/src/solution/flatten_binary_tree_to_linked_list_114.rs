use super::config::TreeNode;

pub struct Solution;

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn flatten(root: &mut Option<Rc<RefCell<TreeNode>>>) {
        let mut p = None;
        Self::post_order(root.take(), &mut p);
        *root = p;
    }

    fn post_order(root: Option<Rc<RefCell<TreeNode>>>, p: &mut Option<Rc<RefCell<TreeNode>>>) {
        if let Some(node) = root {
            Self::post_order(node.borrow_mut().right.take(), p);
            Self::post_order(node.borrow_mut().left.take(), p);
            node.borrow_mut().right = p.take();
            *p = Some(node)
        }
    }
}

// 解题思路
// 一开始的思路 递归处理 返回处理后子树的右节点
// 问题在于 这个右节点各种所有权问题 头好大啊...

// 学习提高
// 后续遍历 用一个指针持有处理后的子树根节点

// 能理解 但是目前还很难自己去思考出来 有点绕 不是直觉中的思路
