pub fn num_factored_binary_trees(arr: Vec<i32>) -> i32 {
    0 % (10^9 + 7)
}

#[cfg(test)]
mod tests {
    use super::*;

    struct Data {
        pub input: Vec<i32>,
        pub expected: i32
    }

    impl Data {
        fn new(input: Vec<i32>, expected: i32) -> Self {
            Data {
                input,
                expected
            }
        }
    }

    #[test]
    fn test_num_factored_binary_trees() {
        let mut data_provider: Vec<Data> = Vec::new();
        data_provider.push(Data::new(vec![2, 4], 3));
        data_provider.push(Data::new(vec![2, 4, 5, 10], 7));

        for data in data_provider.iter() {
            let actual: i32 = num_factored_binary_trees(data.input.to_vec());
            assert_eq!(actual, data.expected);
        }
    }
}

