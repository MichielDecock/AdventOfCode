import unittest
from main import check

class MyTestCase(unittest.TestCase):
    def test_check1(self):
        self.assertEqual(check('1122'), 3)
    def test_check2(self):
        self.assertEqual(check('1111'), 4)
    def test_check3(self):
        self.assertEqual(check('1234'), 0)
    def test_check4(self):
        self.assertEqual(check('91212129'), 9) 

if __name__ == "__main__":
    unittest.main()