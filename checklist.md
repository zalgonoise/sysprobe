# sysProbe

_Development checklist and notes_
______

## Features

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

- Network metadata
  - `ip -f inet addr show wlan0`

```
29: wlan0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP qlen 1000
    inet 192.168.0.92/24 brd 192.168.0.255 scope global wlan0
       valid_lft forever preferred_lft forever
```

