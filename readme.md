# ให้รายชื่อของเนื้อหลายชนิดที่ปะปนกัน

เช่น

```Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.```

ทุกคำเป็นชื่อชนิดเนื้อทั้งหมด โดยที่ไม่ต้องสนใจ , . หรือ space

จงหารายชื่อเนื้อทั้งหมด และระบุจำนวนของเนื้อแต่ละชนิด

ใช้ข้อความจาก <https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text>

- Your project must use Go, Go module, and HTTP framework (GRPC is plus)
- You create a JSON API at this endpoint /beef/summary
- The JSON API must count the number of beef
- We encourage you to write tests, which we will give you some extra score
- We will give you an extra score if you focus on performance.

--- sample response --

```json
{
    "beef": {
        "t-bone": 4,
        "fatback": 1,
        "pastrami": 1,
        "pork": 1,
        "meatloaf": 1,
        "jowl": 1,
        "enim": 1,
        "bresaola": 1
    }
}
```