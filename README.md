# Go-MandelbrotSet
A go package to find the numbers in the mandelbrot set with an additional python file to display them

This uses FindMandelbrotSet.go to iterate through the mandelbrot function 'Z(n+1) = Z(n)^2 + C' 
- uses a max depth of 200 
- outputs a file 'data.csv' with every pair of co-ordinates that fall within the set 
- within is defined as after 200 iterations its <2

Then displayMandelB.py takes this data.csv and outputs it in a scatter graph
- Images folder is for any saved images of the scatter graph
