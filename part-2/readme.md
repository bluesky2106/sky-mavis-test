# Object detection by YOLO

## Document

The final document of the design architecture is `Object-Detection-App.docx` in the same directory of this readme.md file.

## Simple app programmed by python

Beside the document, I also programmed a simple UI by python for Object detection by YOLO network. You can follow the followings instructions for installing dependences, downloading pretrained models and running the app in linux / macos machines.

### Download YOLO pretrained model

Please access to following link and download 3 files, then put them in the directory `source_code/yolo-coco`. Since there is 1 file which is really big, I cannot push them into github.

https://drive.google.com/drive/folders/1wVL5pVBMvd-lfCxMheSxJP9bH16aOk7k?usp=sharing

### Install dependences

If you are using linux OS, access to directory `source_code` and run command

```
bash installation_linux.sh
```

If you are using Mac OS, access to directory `source_code` and run command

```
bash installation_macos.sh
```

### Run the App in local machine

After installing dependencies successfully, you can run the app by typing command:

```
python main.py
```

There is also a `demo.mp4` showing how to use the app.