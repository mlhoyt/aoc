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

fn parse(input: &str) -> Result<Data, String> {
    let vs: Vec<Vec<_>> = input
        .lines()
        .map(|v| {
            v.split_whitespace()
                .map(|v| v.parse::<usize>().expect("cannot parse as usize"))
                .collect()
        })
        .collect();

    // Transpose
    let vs = vs.into_iter().fold(vec![vec![], vec![]], |mut acc, r| {
        acc[0].push(r[0]);
        acc[1].push(r[1]);
        acc
    });

    // Sort
    let vs = vs
        .into_iter()
        .map(|mut v| {
            v.sort();
            v
        })
        .collect();

    return Ok(vs);
}

type Data = Vec<Vec<usize>>;

fn part1(data: &Data) -> usize {
    let c0 = data
        .get(0)
        .expect("could not get first colum of data")
        .into_iter();
    let c1 = data
        .get(1)
        .expect("could not get second colum of data")
        .into_iter();

    c0.zip(c1).map(|(v0, v1)| v0.abs_diff(*v1)).sum()
}

fn part2(data: &Data) -> usize {
    let c0 = data
        .get(0)
        .expect("could not get first colum of data")
        .into_iter();
    let c1 = data
        .get(1)
        .expect("could not get second colum of data")
        .into_iter();

    let c1map: std::collections::HashMap<usize, usize> =
        c1.fold(std::collections::HashMap::new(), |mut acc, v| {
            *acc.entry(*v).or_insert(0) += 1;
            acc
        });

    c0.map(|v| *v * c1map.get(v).map_or(0, |v| *v)).sum()
}
