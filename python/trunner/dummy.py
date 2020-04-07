from trunner.base import TRunner


class DummyRunner(TRunner):
    def __init__(self):
        self.name = "dummy_runner"

    def run(self):
        print("%s is running" % self.name)
        return None


if __name__ == '__main__':
    r = DummyRunner()
    assert r.run() is None
