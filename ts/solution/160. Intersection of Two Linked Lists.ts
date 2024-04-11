function getIntersectionNode(headA: ListNode | null, headB: ListNode | null): ListNode | null {
    while (headA != null ){
        let cur = headB;
        while (cur!=null ){
            if (cur == headA){
                return cur;
            }
            cur = cur.next;
        }
        headA = headA.next;
    }
    return null;
};

function getIntersectionNode2(headA: ListNode|null, headB:ListNode|null):ListNode|null{
    if (headA==null || headB==null) return null;
    let p0 = headA ,p1 = headB;
    while (p0!==p1){
        p0 = p0 && p0.next || headB;
        p1 = p1 && p1.next ||headA;
    }
    return p0;
}

// 简单的暴力破解
// 双指针解法:
// 利用相同路程这一点 使用双指针交叉遍历 如果有交点则必定在交点处相遇 否则两个指针最后将走到两个链表的尾部
