# Buff Duration Formulas

| Buff Duration Formula ID | Buff Duration Calculation Method.  Resulting values represent duration time in "ticks".  |
| :--- | :--- |
| 0 | Not a Buff |
| | **For the following, if the specified buffduration is less than the calculated value, buffduration will be used instead, meaning buffduration serves a cap rather than a raising the duration to a minimum amount of time. <br>"Level" refers to the level of the caster. <br>Be sure to account for parentheses (or lack of parentheses). What you see below is exactly as coded.** |
| 1 | level / 2 |
| 2 | If level > 3, then level / 2 + 5, otherwise 6 |
| 3 | 30 * level |
| 4 | 50 |
| 5 | 2 |
| 6 | level / 2 + 2 |
| 7 | level |
| 8 | level + 10 |
| 9 | 2 * level + 10 |
| 10 | 3 * level + 10 |
| 11 | 30 * (level + 3) |
| 12 | If level > 7, then level / 4, otherwise 1 |
| 13 | 4 * level + 10  |
| 14 | 5 * (level + 2)  |
| 15 | 10 * (level + 10) |
| Any other formula under 200 | Ignored | 
| Any formula 200 or above | If duration > formula, then formula, otherwise duration. | 
| 3600 | Duration if not 0, else 3600 |
| | **The following are fixed and do not use any calculations.** | 
| 50 | 5 Days |
| 51 | Permanent |


# Making Fixed-Length Buffs
If your desired buffduration is 200 ticks or longer, set buffduration to 0 and put the tick value in the buffdurationformula column instead.
If your desired buffduration is under 200, set the buffdurationformula to 200 and the buffduration to your desired tick value.

For your convenience, here are some standard tick values.

| Duration   | Ticks |
|------------|-------|
| 5 minutes  | 50    |
| 10 minutes | 100   |
| 15 minutes | 150   |
| 30 minutes | 300   |
| 45 minutes | 450   |
| 60 minutes | 600   |
| 90 minutes | 900   |
| 2 hours    | 1200  |
| 3 hours    | 1800  |
| 4 hours    | 2400  |
| 5 hours    | 3000  |
| 6 hours    | 3600  |
