#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { val, next: None }
    }

    pub fn create(list: Vec<i32>) -> Option<Box<ListNode>> {
        let mut head = Box::new(Self::new(0));
        // let mut head = Self::new(0);
        let mut tail = &mut head;
        // list.reverse();
        for i in list {
            tail.next = Some(Box::new(ListNode::new(i)));
            tail = tail.next.as_mut().unwrap();
        }
        head.next
    }
}
