import warnings
warnings.filterwarnings("ignore")  # <- put this at the top

import sys
import pickle
import pandas as pd
import numpy as np
from sklearn.preprocessing import StandardScaler

if len(sys.argv) != 9:  # Script name + 8 features
    print("Error: Expected 8 numeric arguments.")
    sys.exit(1)

# Column names must match training data
values = []
# Convert CLI args into a single row DataFrame
for idx, v in enumerate(sys.argv[1:]):
    if idx == 5 or idx == 6:
        values.append(float(v))
    else:
        values.append(int(v))    

df = np.asarray([values]).reshape(1, -1)

with open("scaler.pkl", "rb") as s:
    scalar = pickle.load(s)

input = scalar.transform(df)

# Load model
with open("model.pkl", "rb") as f:
    model = pickle.load(f)

# Predict
prediction = model.predict(input)

# Output result for Go to capture
if prediction[0] == 1:
    print("yes")
else:
    print("no")
