# Receiptify


To run this locally you must have tesseract and all it's dependencies installed

Otherwise:
```docker build . -t receiptify```

```docker run -p 8086:8086 receiptify```

You can then send a base64 encoded image string to the endpoint and get back the text!
