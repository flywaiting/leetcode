use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut res = Box::new(ListNode::new(-1));
        let mut head = head;

        // let mut cur = &mut res.next;
        let mut cur = &mut res;
        while let Some(mut node) = head {
            head = node.next.take();
            // cur = &mut cur.insert(node).next;
            if head.is_none() || head.as_ref()?.val != node.val {
                cur = cur.next.insert(node);
            } else {
                while head.is_some() && head.as_ref()?.val == node.val {
                    head = head.unwrap().next;
                }
            }
        }
        res.next
    }
}

// 解题思路
// 思路很简单 也很清晰 就是用`rust`实现 真的很蛋疼
// 首先 尽量不用有过多的指针 各种引用关系 是真的很操蛋
// 其次 在完成编程前 不要过于相信错误提示 有时候很无脑。。。
// 最后 尽量简洁 `rust`真的会逼着人去思考 还是引用的问题 特别是发生引用、借用的前后顺序

// 暂时吐槽这些
