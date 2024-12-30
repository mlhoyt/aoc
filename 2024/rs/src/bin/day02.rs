fn main() {
    let input = aoc::read_stdin().expect("cannot read stdin");
    // println!("{input}");

    let data = parse(&input).expect("cannot parse input");
    // println!("{data:?}");

    let result1 = part1(&data);
    println!("part1: {result1}");

    let result2 = part2(&data);
    println!("part2: {result2}");
}

type Data = Vec<Vec<usize>>;

fn parse(input: &str) -> Result<Data, String> {
    Ok(input
        .lines()
        .map(|vs| {
            vs.split_whitespace()
                .map(|v| v.parse::<usize>().expect("cannot parse as usize"))
                .collect()
        })
        .collect())
}

fn part1(data: &Data) -> usize {
    data.iter()
        .map(|vs| is_safe_seq(vs))
        .filter(|v| *v == true)
        .count()
}

fn part2(data: &Data) -> usize {
    data.iter()
        .map(|vs| {
            // Produce all permutations of vs with one element removed
            let vsps: Vec<_> = (0..vs.len())
                .into_iter()
                .map(|i| {
                    let mut vsp = vs.clone();
                    vsp.remove(i);

                    vsp
                })
                .collect();

            // Yield if any of the permuations is safe
            vsps.iter().any(|vs| is_safe_seq(vs))
        })
        .filter(|v| *v == true)
        .count()
}

fn is_safe_seq(vs: &[usize]) -> bool {
    // Produce a list of differences between adjacent values
    let vsp: Vec<_> = vs
        .windows(2)
        .map(|w| w[1] as isize - w[0] as isize)
        .collect();

    let is_all_increasing = |vs: &[isize]| -> bool { vs.iter().all(|v| *v > 0) };
    let is_all_decreasing = |vs: &[isize]| -> bool { vs.iter().all(|v| *v < 0) };
    let is_all_diff_gte_1 = |vs: &[isize]| -> bool { vs.iter().all(|v| v.abs() >= 1) };
    let is_all_diff_lte_3 = |vs: &[isize]| -> bool { vs.iter().all(|v| v.abs() <= 3) };

    // Yield if the list of differences satisfies the safe criteria
    (is_all_increasing(&vsp) || is_all_decreasing(&vsp))
        && is_all_diff_gte_1(&vsp)
        && is_all_diff_lte_3(&vsp)
}
