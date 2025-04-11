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
parse = Right . sort . Data.List.transpose . map (map toInt . words) . lines
  where
    toInt x = read x :: Int
    sort = map Data.List.sort

type Data = [[Int]]

part1 :: Data -> Maybe Int
part1 (r : rs) =
    Just $ sum values
  where
    pairs = zip r (head rs)
    values = map (abs . uncurry (-)) pairs

part2 :: Data -> Maybe Int
part2 (r : rs) =
    Just $ sum values
  where
    freqMap = foldr (\x acc -> M.insertWith (+) x 1 acc) M.empty (head rs)
    values = map (\x -> x * M.findWithDefault 0 x freqMap) r
