# How to contribute to the project

This project welcomes any change, improvement, or suggestion!

If you'd like to help its development feel free to open a new issue and raise a pull request.

## IMPORTANT

If you'd like to work on an existing issue, kindly **ask** for it to be assigned to you.

Do you have any struggles with the issue you are working on? Feel free to **tag [me](https://github.com/chasefleming)** in it _and/or_ open a draft pull request.


### How do I make a contribution

If you've never made an open source contribution before or are curious about how contributions operate in our project? Here's a quick rundown!

#### Fork this repository

Fork this repository by clicking on the fork button on the top of [this](https://github.com/chasefleming/elem-go) page.
This will create a copy of this repository in your account `<your-GitHub-username>/<repository-name>`.

#### Clone the repository

Now clone the forked repository to your machine. Go to your GitHub account, open the forked repository, and copy the link provided under `HTTPS` when you click on the green button labeled `code` on the repository page

Open a terminal and run the following git command:

```
git clone "url you just copied"
```

where "URL you just copied" (without quotation marks) is the URL to this repository (your fork of this project). 

For example:

```
git clone https://github.com/username/chasefleming/elem-go.git
```

#### Create a new branch for your changes or fix 

    ```sh
     $ git checkout -b <branch-name>
    ```

#### Setup the project in your local by following the steps listed in the [README.md](https://github.com/chasefleming/elem-go/blob/main/README.md) file

#### Open the project in a code editor and begin working on it
#### Add the contents of the changed files to the "snapshot" git uses to manage the state of the project, also known as the index

    ```sh
    $ git add .
    ```

#### Add a descriptive commit message

    ```sh
    $ git commit -m "Insert a short message of the changes made here"
    ```

#### Push the changes to the remote repository

    ```sh
    $ git push -u origin <branch-name>
    ```

#### Submit a pull request to the upstream repository
