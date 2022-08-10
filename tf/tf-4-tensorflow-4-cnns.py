# https://developers.google.com/codelabs/tensorflow-4-cnns

import tensorflow as tf
print(tf.__version__)
mnist = tf.keras.datasets.fashion_mnist


# reshape
(training_images, training_labels), (test_images, test_labels) = mnist.load_data()
training_images = training_images.reshape(60000, 28, 28, 1)
training_images = training_images / 255.0
test_images = test_images.reshape(10000, 28, 28, 1)
test_images = test_images / 255.0


model = tf.keras.models.Sequential([
    # Add convolution
    tf.keras.layers.Conv2D(64, (2, 2), activation='relu',
                           input_shape=(28, 28, 1)),
    tf.keras.layers.MaxPooling2D(2, 2),

    # Add another convolution
    tf.keras.layers.Conv2D(64, (2, 2), activation='relu'),
    tf.keras.layers.MaxPooling2D(2, 2),

    # Now flatten the output. After this you'll just have the same DNN structure as the non convolutional version
    tf.keras.layers.Flatten(),

    # The same 128 dense layers, and 10 output layers as in the pre-convolution example:
    tf.keras.layers.Dense(128, activation='relu'),
    tf.keras.layers.Dense(10, activation='softmax')
])
model.compile(optimizer='adam',
              loss='sparse_categorical_crossentropy', metrics=['accuracy'])
model.summary()
model.fit(training_images, training_labels, epochs=3)
test_loss, test_accuracy = model.evaluate(test_images, test_labels)
print('Test loss: {}, Test accuracy: {}'.format(test_loss, test_accuracy*100))
