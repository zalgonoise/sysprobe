# sysProbe

_Development checklist and notes_
______

## Response message JSON schema

```json

{
  "net": {
    "sys": {
      "device": "string",
      "id": 0,
      "ipv4": "string",
      "mask": "string"
    },
    "ping": {
      "target": "string",
      "alive": [
        {
          "addr": "string",
          "rtt": 0
        }
      ]
    },
    "ports": [
      {
        "target": "string",
        "proto": "string",
        "ports": [
          0
        ]
      }
    ]
  },
  "power": {
    "status": "string",
    "health": "string",
    "capacity": 0,
    "temp": {
      "int": 0,
      "ext": 0
    },
    "source": "string"
  },
  "timestamp": 0
}
```


## Features

- Network metadata 

  - `ip -f inet addr show wlan0`

```
29: wlan0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP qlen 1000
    inet 192.168.0.92/24 brd 192.168.0.255 scope global wlan0
       valid_lft forever preferred_lft forever
```

  - Ping \*.\*.\*.0/24 networks and list responding hosts

  - Port scan the alive hosts from the ping scan

- System metadata
  
  - Battery metadata 

    - `cat /sys/class/power_supply/battery/uevent`

```
POWER_SUPPLY_NAME=battery
POWER_SUPPLY_STATUS=Discharging
POWER_SUPPLY_CHARGE_TYPE=N/A
POWER_SUPPLY_HEALTH=Good
POWER_SUPPLY_PRESENT=1
POWER_SUPPLY_ONLINE=1
POWER_SUPPLY_TECHNOLOGY=Li-ion
POWER_SUPPLY_VOLTAGE_NOW=4246000
POWER_SUPPLY_VOLTAGE_AVG=4246000
POWER_SUPPLY_CURRENT_NOW=-191
POWER_SUPPLY_CURRENT_AVG=-191
POWER_SUPPLY_CHARGE_NOW=0
POWER_SUPPLY_CAPACITY=98
POWER_SUPPLY_TEMP=310
POWER_SUPPLY_TEMP_AMBIENT=288
POWER_SUPPLY_CHARGE_CONTROL_LIMIT=0
```

  - Termux fallback (if there is no root access):
  
    - `termux-battery-status`


## To do

  - RAM metadata

    - `free --si -hw`

```
              total        used        free      shared     buffers       cache   available
Mem:           1.4G        883M         91M        9.0M         10M        397M        482M
Swap:            0B          0B          0B
```

  - CPU utilization metadata

    - `iostat -zm` 

```
Linux 3.10.49-8935060 (localhost)       01/21/21        _armv7l_        (4 CPU)

avg-cpu:  %user   %nice %system %iowait  %steal   %idle
          25.40    0.30    8.44    0.06    0.00   65.79

Device:            tps    MB_read/s    MB_wrtn/s    MB_read    MB_wrtn
mmcblk0           1.07         0.01         0.01      45048      65101
mmcblk0p1         0.00         0.00         0.00          8          0
mmcblk0p2         0.00         0.00         0.00         48          0
mmcblk0p9         0.00         0.00         0.00          0          0
mmcblk0p12        0.00         0.00         0.00          0          0
mmcblk0p13        0.00         0.00         0.00          3         10
mmcblk0p14        0.00         0.00         0.00          1         28
mmcblk0p15        0.00         0.00         0.00          0         29
mmcblk0p22        0.00         0.00         0.00          0          0
mmcblk0p23        0.00         0.00         0.00          1          0
mmcblk0p24        0.00         0.00         0.00          0          0
mmcblk0p25        0.09         0.00         0.00      15630          0
mmcblk0p26        0.00         0.00         0.00          0        192
mmcblk0p28        0.84         0.01         0.01      29353      64840

```
    
  - System uptime and average load

    - `uptime`

```
20:37:17 up 56 days, 23:40,  load average: 5.15, 5.20, 5.70
```


  - Processes metadata (*nix and app-specific)

    - `tsudo ps -aux | wc -l` (returns a number (_int_))

    - `tsudo ps -aux | grep '/data/data/com.termux/' | wc -l` (returns a number (_int_))



  - Disk / IO statistical metadata

    - Number of open files (*nix and app-specific)

      - `tsudo lsof | wc -l` (returns a number (_int_))

      - `tsudo lsof | grep '/data/data/com.termux/' | wc -l` (returns a number (_int_))

    - I/O Reads/Writes 

      - `iostat -zm` (above)