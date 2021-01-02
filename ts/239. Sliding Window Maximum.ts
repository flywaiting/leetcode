function maxSlidingWindow(nums: number[], k: number): number[] {
    if (k == 1) return nums;

    let que: number[] = [];
    let ans: number[] = [];
    for (let i = 0, n = nums.length; i < n; i++) {
        while (que.length > 0 && nums[que[que.length - 1]] < nums[i]) {
            que.pop();
        }
        if (que.length && que[0] < i - k + 1) {
            que.shift();
        }
        que.push(i)
        if (i+1>=k) {
            ans.push(nums[que[0]]);
        }
    }
    return ans
};