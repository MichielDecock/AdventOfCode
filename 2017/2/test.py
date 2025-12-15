import unittest
from main import checkSum

class MyTestCase(unittest.TestCase):
    def test_check1(self):
        self.assertEqual(checkSum([[5, 9, 2, 8], [9, 4, 7, 3], [3, 8, 6, 5]]), 9)


if __name__ == "__main__":
    unittest.main()