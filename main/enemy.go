components {
  id: "enemy_actions"
  component: "/main/enemy_actions.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"tank_down\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"tanks\"\n"
  "mask: \"rockets\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 8.805526\n"
  "  data: 11.698438\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
