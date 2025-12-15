import unittest
from main import check

class MyTestCase(unittest.TestCase):
    def test_check1(self):
        self.assertEqual(check('1212'), 6)
    def test_check2(self):
        self.assertEqual(check('1221'), 0)
    def test_check3(self):
        self.assertEqual(check('123425'), 4)
    def test_check4(self):
        self.assertEqual(check('123123'), 12)
    def test_check5(self):
        self.assertEqual(check('12131415'), 4)

if __name__ == "__main__":
    unittest.main()