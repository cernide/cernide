Before we can train our experiments we need to deploy Polyaxon.

1. Change to examples folder, and specifically to cifar10

    ```bash
    $ cd /path/to/polyaxon-examples/tensorflow/cifar10
    ```

2. Deploy Polyaxon

    Add polyaxon charts to helm.

    ```bash
    $ helm repo add polyaxon https://charts.polyaxon.com
    $ helm repo update
    ```

    Now you can install Polyaxon with the `gke/polyaxon-config.yml` file,
    it overrides the default values from Polyaxon, by adding the data and outputs `existingClaim`
    with the values we just created. It also sets the compatible NVidia paths with GPU nodes.

    ```bash
    $ helm install polyaxon/polyaxon --name=polyaxon --namespace=polyaxon -f ../../gke/polyaxon-config.yml
    ```

3. Configure Polyaxon-CLI

    When the deployment is finished, you will receive a list of instructions in order to configure the cli.

4. Login with the superuser created by Polyaxon

    the instructions will tell you the superuser's username and a command to get the password.

    ```bash
    $ polyaxon login --username=root --password=rootpassword
    ```

5. Create a project

    ```bash
    $ polyaxon project create --name=cifar10 --description='Train and evaluate a CIFAR-10 ResNet model on polyaxon'
    ```

6. Init the project

    ```bash
    $ polyaxon init cifar10
    ```

7. Create the first experiment

    Unless you copied the data manually to the data share create in the previous section,
    the first time you create an experiment, it will create the data.

    Start an experiment with the default `polyaxonfile.yml`

    ```bash
    $ polyaxon run -u
    ```

8. Watch the logs

    ```bash
    $ polyaxon experiment -xp 1 logs

      Building -- creating image -
      master.1 -- INFO:tensorflow:Using config: {'_model_dir': '/outputs/root/cifar10/experiments/1', '_save_checkpoints_secs': 600, '_num_ps_replicas': 0, '_keep_checkpoint_max': 5, '_session_config': gpu_options {
      master.1 --   force_gpu_compatible: true
      master.1 -- }
    ```

9. Watch the resources

    When the experiment starts training, you can also watch the logs

    ```bash
    $ polyaxon experiment -xp 1 resources

      Job       Mem Usage / Limit    CPU% - CPUs
    --------  -------------------  ---------------
    master.1  1.26 Gb / 6.79 Gb    120.11% - 6
    ```

    N.B. Azure seems to have an issue with reporting docker cpu resources.

10. Run a distributed experiment

    ```bash
    $ polyaxon run -f polyaxonfile_distributed.yml -u
    ```


8. Watch the logs

    ```bash
    $ polyaxon experiment -xp 2 logs

      Building -- creating image -
      worker.1 -- INFO:tensorflow:image after unit resnet/tower_0/stage_1/residual_v1_3/: (?, 16, 16, 32)
      worker.1 -- INFO:tensorflow:image after unit resnet/tower_0/stage_1/residual_v1_4/: (?, 16, 16, 32)
      worker.1 -- INFO:tensorflow:image after unit resnet/tower_0/stage_1/residual_v1_5/: (?, 16, 16, 32)
      worker.3 -- INFO:tensorflow:Using config: {'_model_dir': '/outputs/root/cifar10/experiments/2', '_save_checkpoints_secs': 600, '_num_ps_replicas': 0, '_keep_checkpoint_max': 5, '_session_config': gpu_options {
      worker.3 --   force_gpu_compatible: true
      worker.1 -- INFO:tensorflow:image after unit resnet/tower_0/stage_1/residual_v1_6/: (?, 16, 16, 32)
      worker.3 -- }
      worker.3 -- allow_soft_placement: true
      worker.3 -- , '_tf_random_seed': None, '_task_type': None, '_environment': 'local', '_is_chief': True, '_cluster_spec': <tensorflow.python.training.server_lib.ClusterSpec object at 0x7fc7e9f53850>, '_tf_config': gpu_options {
      worker.3 --   per_process_gpu_memory_fraction: 1.0
      worker.3 -- }
      worker.3 -- , '_num_worker_replicas': 0, '_task_id': 0, '_save_summary_steps': 100, '_save_checkpoints_steps': None, '_evaluation_master': '', '_keep_checkpoint_every_n_hours': 10000, '_master': '', '_log_step_count_steps': 100}
      master.1 -- INFO:tensorflow:Using config: {'_model_dir': '/outputs/root/cifar10/experiments/2', '_save_checkpoints_secs': 600, '_num_ps_replicas': 0, '_keep_checkpoint_max': 5, '_session_config': gpu_options {
      master.1 --   force_gpu_compatible: true
      master.1 -- }
      ...
      worker.2 -- INFO:tensorflow:Evaluation [2/100]
      worker.4 -- INFO:tensorflow:Evaluation [2/100]
      worker.2 -- INFO:tensorflow:Evaluation [3/100]
      worker.4 -- INFO:tensorflow:Evaluation [3/100]
      worker.2 -- INFO:tensorflow:Evaluation [4/100]
      worker.4 -- INFO:tensorflow:Evaluation [4/100]
      worker.2 -- INFO:tensorflow:Evaluation [5/100]
      worker.4 -- INFO:tensorflow:Evaluation [5/100]
      worker.2 -- INFO:tensorflow:Evaluation [6/100]
      worker.4 -- INFO:tensorflow:Evaluation [6/100]
      worker.2 -- INFO:tensorflow:Evaluation [7/100]
      worker.4 -- INFO:tensorflow:Evaluation [7/100]
      worker.2 -- INFO:tensorflow:Evaluation [8/100]
      worker.4 -- INFO:tensorflow:Evaluation [8/100]
      worker.2 -- INFO:tensorflow:Evaluation [9/100]
      worker.4 -- INFO:tensorflow:Evaluation [9/100]
    ```

9. Watch the resources

    When the experiment starts training, you can also watch the logs

    ```bash
    $ polyaxon experiment -xp 2 resources

    Job       Mem Usage / Total    CPU% - CPUs
    --------  -------------------  -------------
    master.1  1.23 Gb / 55.03 Gb   0.01% - 6
    worker.2  1.1 Gb / 6.79 Gb     73.33% - 2
    worker.3  1.26 Gb / 55.03 Gb   246.32% - 6
    worker.4  1.12 Gb / 6.79 Gb    67.03% - 2
    ps.5      1.21 Gb / 55.03 Gb   272.41% - 6
    ```
