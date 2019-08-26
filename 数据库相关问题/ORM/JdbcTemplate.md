# 入门

## 在Spring Boot种使用JdbcTemplate
### 引入依赖
```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-jdbc</artifactId>
</dependency>
```

### 写代码在Dao层或者Mapper
这里主要演示在Dao的实现类里使用JdbcTemplate，所以其它模块代码的编写就不展示了，具体可参考[源码](https://github.com/wuyouzhuguli/SpringAll/tree/master/04.Spring-Boot-JdbcTemplate)。

StudentDaoImp类代码：
```java
@Repository("studentDao")
public class StudentDaoImp implements StudentDao {

    @Autowired
    private JdbcTemplate jdbcTemplate;
    
    @Override
    public int add(Student student) {
        // String sql = "insert into student(sno,sname,ssex) values(?,?,?)";
        // Object[] args = { student.getSno(), student.getName(), student.getSex() };
        // int[] argTypes = { Types.VARCHAR, Types.VARCHAR, Types.VARCHAR };
        // return this.jdbcTemplate.update(sql, args, argTypes);
        String sql = "insert into student(sno,sname,ssex) values(:sno,:name,:sex)";
        NamedParameterJdbcTemplate npjt = new NamedParameterJdbcTemplate(this.jdbcTemplate.getDataSource());
        return npjt.update(sql, new BeanPropertySqlParameterSource(student));
    }
    
    @Override
    public int update(Student student) {
        String sql = "update student set sname = ?,ssex = ? where sno = ?";
        Object[] args = { student.getName(), student.getSex(), student.getSno() };
        int[] argTypes = { Types.VARCHAR, Types.VARCHAR, Types.VARCHAR };
        return this.jdbcTemplate.update(sql, args, argTypes);
    }
    
    @Override
    public int deleteBysno(String sno) {
        String sql = "delete from student where sno = ?";
        Object[] args = { sno };
        int[] argTypes = { Types.VARCHAR };
        return this.jdbcTemplate.update(sql, args, argTypes);
    }
    
    @Override
    public List<Map<String, Object>> queryStudentsListMap() {
        String sql = "select * from student";
        return this.jdbcTemplate.queryForList(sql);
    }
    
    @Override
    public Student queryStudentBySno(String sno) {
        String sql = "select * from student where sno = ?";
        Object[] args = { sno };
        int[] argTypes = { Types.VARCHAR };
        List<Student> studentList = this.jdbcTemplate.query(sql, args, argTypes, new StudentMapper());
        if (studentList != null && studentList.size() > 0) {
            return studentList.get(0);
        } else {
            return null;
        }
    }
}
```