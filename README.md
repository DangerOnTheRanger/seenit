seenit - a perceptual-hash-based repost checker
===============================================

# What? 

**seenit** is a repost checker. Punch in the name of a community and
upload an image, and it'll tell you if someone else has already done the same.
You'll never have to worry about reposts in the meme chat again! It's also uses
a basic form of perceptual hashing, so it'll handle resizing and things like JPEG
compression fairly well.

# Why?

~~Because I had one too many memes of mine reposted in a particular chat and decided
enough was enough, and needed a project to show off my [neumorphic CSS library](https://github.com/DangerOnTheRanger/morphy-css) anyway.~~ Don't worry about it.

# Screenshots

![Screenshot](https://raw.githubusercontent.com/DangerOnTheRanger/seenit/main/screenshot.png)

# Building

The included `Dockerfile` is the intended way to go about doing things. For persisting
information (seen hashes, etc.), there will be a `database.db` file in the container that
you can mount and save.

# License

**seenit** is licensed under the MIT license.
