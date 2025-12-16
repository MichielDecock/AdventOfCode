import unittest
from main import spiral

class MyTestCase(unittest.TestCase):
    def test_spiral1(self):
        self.assertEqual(spiral(2), 4)
    def test_spiral2(self):
        self.assertEqual(spiral(4), 5)
    def test_spiral3(self):
        self.assertEqual(spiral(5), 10)
    def test_spiral4(self):
        self.assertEqual(spiral(10), 11)
    def test_spiral5(self):
        self.assertEqual(spiral(23), 25)

if __name__ == "__main__":
    unittest.main() 