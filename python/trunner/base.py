class TRunner:

    def upload_model(self):
        raise NotImplementedError

    def upload_tensorboard_event(self):
        raise NotImplementedError

    def upload_evaluate(self):
        raise NotImplementedError

    def upload_inference(self):
        raise NotImplementedError
