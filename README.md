# star-drive

A novel(?) 2-axis cartesian motion system.


## Inspiration

Havig found a pair of high pitch ballscrews on Amazon, the need arose
for a budget friendly linear motion system with corresponding stiffness
for accuracy while being light-weight for speed.
The target applications for the project were engraving (PCB milling) and
3D printing.

A standard X-Y system has to balance cost, stiffness, and speed - a strong
frame is heavy, or both sides of one axis need to be driven by either
independent drive systems or cabling to prevent racking, and the entire mass
of the second axis is supported and moved by the first.

A crossed X-Y system greatly reduces the moving mass (all motors are
stationary), but has problems with racking on both axes, and it requires
additional linear rail.

The resulting design appears to be new, but you should perform your own
exhaustive search before doing anything with it.


## Description

Crossed rails are used to minimize driven mass, but the cross angle is nominally
60 degrees rather than 90.
This allows both cross rails to ride on the same drive rails with the motors
and ballscrews located on the same or opposite sides.

-- (overview diagram here)

Driving both cross rails the same distance along the drive rails causes the
platform to move an equal distance parallel to the drive rails.
Driving one cross rail a distance along the drive rails causes the platform
to move an equal distance along the other cross rail.
To move the platform perpendicular to the drive rails, the cross rails must
each be moved about 4/7 the desired distance, but in opposite directions.

The name, "* drive", was loosely based on the appearance of the crossed rails
and the simple directions of motion.

-- (motion diagrams here)

Compared to the standard 90 cross angle, one set of rails and a motor are
eliminated, as is the tendency for racking.
Calibration for orthogonal movement requires a handful of trivial measurements
with corrections applied in the controller software.

The biggest drawbacks are the cost of the ballscrews and the fact that the
range of usable motion is significantly less than the area enclosed by the
drive rails.

The system may be adapted to a Y-Z motion by extending the one set of drive
rails (Y range) and optionally adjusting the cross angle for finer control
of the Z position.

-- (Y-Z diagram here)


## Prototype

![First prototype](first-prototype.png)


## Inverse Kinematics


## Calibration
