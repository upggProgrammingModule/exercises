# exercises
A set of exercises to introduce genomics and programming in Go (golang)

# Getting Started
1. Make a github.com account
   * Send an email to Craig with your username
2. Setup github email preferences
   * Detailed instructions: https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-user-account/managing-email-preferences/setting-your-commit-email-address
   * Click in the top right and go to "Settings"
   * Go to "Emails"
   * Verify your email address
   * Click "Keep my email address private"
   * Click "Block command line pushes that expose my email"
   * In the text under "Keep my email address private" notice your noreply email address and copy this somewhere for later
3. Make SSH keys
   * Detailed instructions here: https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent
   * You do not need to add this key to your github account, we will be adding it to your Duke OIT account.
4. Add your laptop's public key to your Duke OIT account
   * Add the key you just generated to “SSH Public keys” under “Advanced User Options” at: https://idms-web.oit.duke.edu/portal
5. Log in to the virtual linux machine I have created for the course
   * Go to vcm.duke.edu
   * Reserve a virtual machine
   * Under "Linux Apps" you should see "Ubuntu GoLang"
   * Click on it and VM should be created for you.  It claims this can take 30 minutes, but the times I have tried it has taken about five minutes.  You will get an email when it is ready.
   * Ssh from your laptop to your virtual machine using your netid.  There will be an example command in the email from OIT.
   * If you ssh keys were setup properly, you should not be prompted for a password or two-factor authentication.
6. Enjoy the victory of having ssh'ed into your own virtual machine.  All instructions are now for your virtual machine.
7. You should be able to type "go version" and have it tell you the version of Go (golang) installed.
   * Let me know if this does not work
8. Make an ssh key for this virtual machine using the same steps in 3.
   * Add this ssh key to your github account: https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account
9. Setup git on this virtual machine:
   * git config --global user.name "Your Name"
   * git config --global user.email yourGitHubNoReplyEmailAdressFromEarlier
10. Go used to have a strict way to setup your code.  You no longer need to do things in this strict layout, but we will for the course because it is still the standard and helpfu.
   * In your home directory create a few directories:
   * mkdir -p go/src/github.com/upggProgrammingModule
   * mkdir -p go/pkg
   * mkdir -p go/bin
11. Clone the code repository for this course
   * cd go/src/github.com/upggProgrammingModule
   * git clone git@github.com:upggProgrammingModule/exercises.git
12. Run the test program
   * cd exercises/sayHello
   * go build
   * ./sayHello
13. Find the bugs!
   * cd ../parrot
   * go test
   * create your own branch: git checkout -b yourNameParrotFix
   * edit parrot.go to fix the bugs
   * run go test to make sure you fixed them
   * push your changes to github: git push (follow it's suggested command)

   
