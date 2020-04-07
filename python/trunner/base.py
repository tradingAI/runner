class TRunner:

    def upload_model(self, *args):
        raise NotImplementedError

    def upload_tensorboard_event(self, *args):
        raise NotImplementedError

    def upload_evaluate(self, *args):
        raise NotImplementedError

    def upload_inference(self, *args):
        raise NotImplementedError

    def train(self, *args):
        raise NotImplementedError

    def eval(self, *args):
        raise NotImplementedError

    def inference(self, *args):
        raise NotImplementedError
