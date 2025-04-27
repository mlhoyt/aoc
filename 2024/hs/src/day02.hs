import Data.List
import qualified Data.Map as M
import qualified Data.Maybe

main = do
    input <- getContents -- aka read_stdin
    -- putStrLn ("input: " ++ input)
    case parse input of
        Left err -> do
            putStrLn ("failed parsing input: " ++ err)
        Right v -> do
            -- putStrLn ("data: " ++ show v)
            putStrLn ("part1: " ++ show (part1 v))
            putStrLn ("part2: " ++ show (part2 v))

parse :: String -> Either String Data
parse = Right . map (map toInt . words) . lines
  where
    toInt x = read x :: Int

type Data = [[Int]]

part1 :: Data -> Maybe Int
part1 = Just . length . filter id . map isSafeSeq

part2 :: Data -> Maybe Int
part2 = Just . length . filter id . map (any isSafeSeq . expand)
  where
    expand :: [a] -> [[a]]
    expand xs = zipWith removeAt [0 ..] $ replicate n xs
      where
        n = length xs
        removeAt i xs = take i xs ++ drop (i + 1) xs

isSafeSeq :: (Num a, Ord a) => [a] -> Bool
isSafeSeq xs = (allIncreasing || allDecreasing) && allAbsGte1 && allAbsLte3
  where
    vs = zipWith (-) xs (tail xs)
    allIncreasing = all (> 0) vs
    allDecreasing = all (< 0) vs
    allAbsGte1 = all ((>= 1) . abs) vs
    allAbsLte3 = all ((<= 3) . abs) vs
