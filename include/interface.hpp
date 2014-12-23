#ifndef RSD_INTERFACE_HPP_
#define RSD_INTERFACE_HPP_

#define DECLARE_INTERFACE(interface)\
public:\
  virtual ~interface() = default;\
protected:\
  interface() = default;\
  interface(const interface &) = default;\
  interface &operator =(const interface &) = default;

#endif  // RSD_INTERFACE_HPP_
