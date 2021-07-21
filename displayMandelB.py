import pandas as pd
df  = pd.read_csv("data.csv")
df.plot(kind='scatter',x='x',y='y') # scatter plot
