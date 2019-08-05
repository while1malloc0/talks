# Defeating an evil AI
## with the Linux kernel

<!--
How many of you have applications running on Kubernetes?
How many of you know feel like you really know what a container is?
-->

---

# VMs

<!--
Containers are often presented as "light weight" VMs.
It's good to have an understanding of how they're different.
-->

---

# VMs
#### Virtualize an entire operating system, kernel and all

<!--
VMs create a virtual copy of an entire operating system in software.
This includes the operating system's kernel.
In contrast, containers share the kernel of their host operating system.
-->

---

# What is a kernel?

<!--
Kernels manage the operations of the hardware.
It takes care of things like managing memory, CPU scheduling, and drivers for specific peripherals.
It also presents a user interface for interacting with that hardware. In Linux, that's called libc.
-->

---

[[ Picture of kernel diagram ]]

---

# Let's build a container!

<!--
So now that we understand our basic building blocks, let's build a container.
The first thing we need to build a container is a process to be contained.
Let's take a look at the app that we're going to be containing.
-->

---

[[ Picture of Ultron ]]

<!--
Our process is going to be an evil AI.
Here are the rules: the evil AI app boots up, and from there can run any command with root permissions.
The evil AI app is going to attempt to run a command that will do the most damage to the machine that it's running on as possible.

We boot up our app and it runs this command
-->

---

# rf -rf /

<!--
Ok, so that's clearly not good.
Our evil AI has just wiped our root folder.
We can't keep our evil AI from running this command, but maybe we can trick it into thinking / is a different folder.
-->

---

# chroot

<!--
This is where chroot comes in, and it's the first tool in our container toolkit.
chroot tricks a process into thinking that its root folder is a different folder than it actually is.
-->

---

[[ chroot demo ]]

<!--
Ok, so at worst our evil AI can just delete itself.
But there's a problem here. Our app can't run at all because it's missing libraries.
-->

---

# Minimal file systems

<!--
[[ Some words about minimal file systems ]]
Ok, so now our evil AI app runs, and can only delete itself.
Let's see what it does this time.
-->

---

# ps -e -o pid | xargs kill -9

<!--
Ouch. For those that don't spend their time thinking up malicious bash commands, this will kill every process running on the box.
Again, we can't keep the evil AI from running the command, but maybe we can trick it into not seeing any processes.
-->

---

# Namespaces

<!--
[[ Words about namespaces ]]
Ok, so now our evil AI can't mess with other processes.
What does it do next?
-->

---

# while true; do yes > /dev/null &; done

<!--
This one's a little less direct.
Instead of interfering with processes directly, our evil AI is going to claim all of the CPU for itself, grinding the rest of our processes to a halt.
-->

---

# CGroups

<!--
[[ Words about cgroups ]]
Now our evil AI will be throttled when it uses more than its fair share of CPU.
What does it try now?
-->

---

# shutdown -r now

<!--
Huh, well, ok.
Now our evil AI is just going to shut down the computer every time it boots.
At this point, we're a bit stuck. We can't stop the AI from running this command as root.
However, we have one more trick up our sleeve: we can remove permissions from the root user.
-->

---

# Capabilities

<!--
[[ Words about capabilities ]]
-->

---

# Docker

<!--
So how does docker play into all of this?
-->

---

# Docker
#### Gives us a packaging format for minimal file systems

<!--
-->

---

# Docker
#### Coordinates chroot, cgroups, namespaces, and capabilities

<!--
-->

---

# Docker
#### Coordinates shared file systems between host and container

<!--
-->

---

# Kubernetes
#### Coordinates groups of Docker containers

<!--
-->

---

# Thanks!

<!--
-->
