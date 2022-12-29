import cv2

print("[*] Stating")

def readNames():
    classNames = []
    classFile = 'coco.names'

    with open(classFile, 'rt') as f :
        classNames = f.read().rstrip('\n').split('\n')
    
    return classNames

img = cv2.imread('template_image.jpeg')

configPath = 'ssd_mobilenet_v3_large_coco_2020_01_14.pbtx'
weightsPath = 'frozen_inference_graph.pb'

net = cv2.dnn_DetectionModel(weightsPath, configPath)
net.setInputSize(320, 320)
net.setInputScale(1.0 / 127.5)
net.setInputMean((127.5, 127.5, 127.5))
net.setInputSwapRB(True)

classIds, confidences, boundingBoxes = net.detect(img, confThreshold=0.5)
print(classIds, boundingBoxes)

for classId, confidence, box in zip(classIds.flatten(), confidences.flatten(), boundingBoxes):
    cv2.rectangle(img, box, color=(0, 255, 0))



cv2.imshow("Output", img)
cv2.waitKey(0)