# -*- coding:utf-8 -*-
import torch
from torch.multiprocessing import set_start_method

import tbase
import tenvs
from tbase.common.cmd_util import (common_arg_parser, make_env, make_eval_env,
                                   make_infer_env, set_global_seeds)
from tbase.common.logger import logger
from tbase.run import get_agent
from trunner.base import TRunner


class TbaseRunner(TRunner):
    def __init__(self):
        self.name = "tbase_runner"
        logger.info("tbase version: %s" % tenvs.__version__)
        logger.info("tbase version: %s" % tbase.__version__)

    def run(self):
        logger.info("%s is running" % self.name)
        args = common_arg_parser()
        if args.debug:
            import logging
            logger.setLevel(logging.DEBUG)
        set_global_seeds(args.seed)
        logger.info("tbase.run set global_seeds: %s" % str(args.seed))
        if torch.cuda.is_available():
            if args.num_env > 1 and args.device != 'cpu':
                set_start_method('spawn')
        env = make_env(args=args)
        print("\n" + "*" * 80)
        logger.info("Initializing agent by parameters:")
        logger.info(str(args))
        agent = get_agent(env, args)
        if not args.eval and not args.infer:
            logger.info("Training agent")
            agent.learn()
            logger.info("Finished, tensorboard --logdir=%s" %
                        args.tensorboard_dir)
        # eval models
        if args.eval:
            eval_env = make_eval_env(args=args)
            agent.eval(eval_env, args)

        # infer actions
        if args.infer:
            infer_env = make_infer_env(args=args)
            agent.infer(infer_env)


if __name__ == '__main__':
    r = TbaseRunner()
    r.run()
