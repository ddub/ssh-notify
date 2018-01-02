ssh-notify
==========

Send Slack notifications when a user accesses a secure shell of a node,
so that there are no surprises from manual changes made after node provisioning.

Based upon similar work for [google compute engine instances](https://cloud.google.com/community/tutorials/send-connect-notification-to-slack-from-google-compute-engine)

# Install

1. create an incomming webhook in your slack team
1. edit example/gke.yaml to change the webhook to the one created above
1. kubectl apply -f example/gke.yaml

Now you can ssh into one of the nodes and a slack notification will be sent!

# Process
The example/gke.yaml provisions a daemonset to provision a pod onto each of your nodes.
This pod runs a docker container that has access to /etc/pam.d and /var/lib/toolbox/ssh-notify on the host.

The docker container will 
1. Copy the notify binary into the toolbox directory
1. Write the config.yaml file into that directory with your slack webhook address
1. Ensures that /etc/pam.d/system-remote-login will call the notify binary with the configuration when a user logs in.

Then it will re-check every 60 seconds that this is still valid.
