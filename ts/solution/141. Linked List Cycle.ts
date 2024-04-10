function hasCycle(head: ListNode | null): boolean {
    let tag = 1e7;
    while (head != null ){
        if (head.val == tag) {
            return true;
        }
        head.val = tag;
        head = head.next;
    }
    return false;
};

function hasCycle2(head:ListNode |null):boolean {
    let fast = head?.next?.next;
    let slow = head?.next;
    while( fast != null ){
        if (fast == slow) return true;
        fast = fast?.next?.next;
        slow = slow?.next;
    }
    return false;
}

// 解题思路
// 特殊标记 直接暴力爆破 缺点就是会破坏原始数据

// 快慢指针
// Floyd 判圈算法, 通过快慢指针的倍数关系 只要存在环 就一定能相遇
// 此外 还能求解环长度 这个显而易见; 环的起始点, 还是利用了倍数关系
