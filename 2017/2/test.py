import unittest
from main import checkSum

class MyTestCase(unittest.TestCase):
    def test_check1(self):
        self.assertEqual(checkSum([[5, 1, 9, 5], [7, 5, 3], [2, 4, 6, 8]]), 18)


if __name__ == "__main__":
    unittest.main()