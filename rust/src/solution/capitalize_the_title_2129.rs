impl Solution {
    pub fn capitalize_title(title: String) -> String {
        title
            .split_whitespace()
            .map(|word| {
                if word.len() < 3 {
                    word.to_lowercase()
                } else {
                    word[..1].to_uppercase() + &word[1..].to_lowercase()
                }
            })
            .collect::<Vec<String>>()
            .join(" ")
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    #[test]
    fn test_2129() {
        assert_eq!(
            super::Solution::capitalize_title("capiTalIze tHe titLe".to_string()),
            "Capitalize The Title"
        );
        assert_eq!(
            super::Solution::capitalize_title("First leTTeR of EACH Word".to_string()),
            "First Letter of Each Word",
        );
        assert_eq!(
            super::Solution::capitalize_title("i lOve leetcode".to_string()),
            "i Love Leetcode"
        );
    }
}
