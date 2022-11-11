# Polyweather

A tiny golang program to get current weather observation data from the weather.gov API  (Probably US-only).  Meant to be used with polybar / waybar.

# Usage

[Fetch your station ID](https://forecast.weather.gov/stations.php) you'd like to use for realtime updates beforehand.

`./polywather <station ID>`

The output will be a string that can be fed into polybar!
