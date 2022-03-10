use crate::AocError;

// Leverage https://github.com/mlhoyt/aoc2020/blob/main/rs/src/bin/day11part2.rs "Layout"
// 2D grid abstraction.

#[derive(Debug)]
pub struct Grid2D<T: Copy> {
    grid: Vec<T>,
    width: usize,
    height: usize,
}

impl<T: Copy> Grid2D<T> {
    pub fn new(rows: &[Vec<T>]) -> Result<Self, AocError> {
        let mut rv = Self {
            grid: vec![],
            width: 0,
            height: 0,
        };

        for (i, r) in rows.iter().enumerate() {
            if i == 0 {
                rv.width = r.len();
            } else if rv.width != r.len() {
                return Err(AocError::new(
                    format!(
                        "row {} has length {} which does not match the previous length {}",
                        i,
                        r.len(),
                        rv.width
                    )
                    .as_str(),
                ));
            }

            r.iter().for_each(|v| rv.grid.push(*v));

            rv.height += 1;
        }

        Ok(rv)
    }

    pub fn get_yx(&self, y: usize, x: usize) -> Option<T> {
        match self.yx_to_index(y, x) {
            None => None,
            Some(i) => Some(self.grid[i]),
        }
    }

    fn index_to_yx(&self, i: usize) -> (usize, usize) {
        let row = (i / self.width) as usize;
        let col = (i % self.width) as usize;

        (row, col)
    }

    fn yx_to_index(&self, y: usize, x: usize) -> Option<usize> {
        if y < (self.height) && x < (self.width) {
            let n = (y * self.width) + (x);
            Some(n)
        } else {
            None
        }
    }

    pub fn iter(&self) -> Grid2DIter<T> {
        Grid2DIter::<T> {
            grid: self,
            index: 0,
        }
    }
}

pub struct Grid2DIter<'a, T: Copy> {
    grid: &'a Grid2D<T>,
    index: usize,
}

impl<'a, T: Copy> Iterator for Grid2DIter<'a, T> {
    type Item = Grid2DPoint<T>;

    fn next(&mut self) -> Option<Self::Item> {
        if self.index >= self.grid.grid.len() {
            None
        } else {
            let pos = self.grid.index_to_yx(self.index);
            self.index += 1;

            Some(Self::Item {
                x: pos.1,
                y: pos.0,
                value: self.grid.get_yx(pos.0, pos.1).unwrap(),
            })
        }
    }
}

#[derive(Hash, PartialEq, Eq, Clone)]
pub struct Grid2DPoint<T> {
    pub x: usize,
    pub y: usize,
    pub value: T,
}
