from trunner.base import TRunner


class DummyRunner(TRunner):
    def __init__(self):
        self.name = "dummy_runner"

    def upload_model(self):
        print("uploading model")
        return None

    def upload_tensorboard_event(self):
        print("uploading tensorboard event")
        return None

    def upload_evaluate(self):
        print("uploading evaluate results")
        return None

    def upload_inference(self):
        print("uploading inference results")
        return None


if __name__ == '__main__':
    r = DummyRunner()
    assert r.upload_model() is None
    assert r.upload_tensorboard_event() is None
    assert r.upload_evaluate() is None
    assert r.upload_inference() is None
