import unittest
from main import escape

class MyTestCase(unittest.TestCase):
    def test_escape1(self):
        self.assertEqual(escape([0, 3, 0, 1, -3]) ,5)

if __name__ == '__main__':
    unittest.main()
