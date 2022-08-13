use std::collections::HashMap;

pub fn num_factored_binary_trees(arr: Vec<i32>) -> i32 {
    let mut result = 0;
    let modulo = 1_000_000_007;

    if is_valid_input(&arr) {
        let mut numbers = arr.clone();
        let mut hash: HashMap<i32, i32> = HashMap::new();

        numbers.sort();

        for (i, num) in numbers.iter().enumerate() {
            hash.insert(*num, 1);
            let mut j = 0;
            while j < i {
                let first = numbers[j];
                if num % first == 0 {
                    let mut count = *hash.get(num).unwrap();
                    let candidate_count = hash.get(&first).unwrap();
                    let second = num / first;
                    let mult = hash.get(&second).unwrap_or(&0);

                    count = count + candidate_count * mult;
                    count = count % modulo;
                    hash.insert(*num, count);
                }
                j = j + 1;
            }
        }

        for (_, count) in hash.iter() {
            result = result + count;
        }
    }
    result % modulo
}

fn is_valid_input(arr: &Vec<i32>) -> bool {
    arr.len() > 0 && is_positive(arr) && is_unique(arr)
}

fn is_positive(arr: &Vec<i32>) -> bool {
    arr.iter().find(|&&x| x < 1) == None
}

fn is_unique(arr: &Vec<i32>) -> bool {
    let mut hash: HashMap<i32, i32> = HashMap::new();
    for num in arr.iter() {
        if hash.contains_key(num) {
            return false;
        }
        hash.insert(*num, 1);
    }
    true
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
        data_provider.push(Data::new(vec![15, 13, 22, 7, 11], 5));

        for data in data_provider.iter() {
            let actual: i32 = num_factored_binary_trees(data.input.to_vec());
            assert_eq!(actual, data.expected);
        }
    }
}

