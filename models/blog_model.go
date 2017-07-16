package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Blog struct {
	ID        int64 `orm:"column(id);pk;auto"`
	Title     string
	Category  *Category `orm:"rel(fk)"`
	Content   string    `orm:"type(text)"`
	Views     int32
	StatusId  int8      /*0:draft, 1:released*/
	Type      int8      `orm:"null"` /*0:original, 1:translation, 2:reprint*/
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (this *Blog) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func GetAllBlogs() (blogs []*Blog) {
	o := orm.NewOrm()
	o.QueryTable("blog").RelatedSel().All(&blogs)
	return blogs
}

func DeleteBlog(id int64) error {
	o := orm.NewOrm()
	blog := Blog{ID: id}
	_, err := o.Delete(&blog)
	return err
}

func UpdateBlogStatus(id int64) error {
	o := orm.NewOrm()
	blog := Blog{ID: id}
	var err error
	if o.Read(&blog) == nil {
		if blog.StatusId == 1 {
			blog.StatusId = 0
		} else if blog.StatusId == 0 {
			blog.StatusId = 1
		}
		_, err = o.Update(&blog)
	}

	return err
}

func AddBlog(title string, categoryId int64, content string) error {
	o := orm.NewOrm()
	catetory := Category{ID: categoryId}
	blog := Blog{
		Title:    title,
		Category: &catetory,
		Content:  content}
	_, err := o.Insert(&blog)

	return err
}

func EditBlog(id int64, title string, categoryId int64, content string) error {
	o := orm.NewOrm()
	blog := Blog{ID: id}
	var err error
	if o.Read(&blog) == nil {
		blog.Title = title
		category := Category{ID: categoryId}
		blog.Category = &category
		blog.Content = content
		_, err = o.Update(&blog)
	}
	return err
}
func GetBlog(id int64) *Blog {
	o := orm.NewOrm()
	blog := &Blog{}
	o.QueryTable("blog").Filter("ID", id).RelatedSel().One(blog)
	return blog
}

func GetDashBoardBlogs() ([]string, map[string][]*Blog) {
	o := orm.NewOrm()
	qs := o.QueryTable("blog")
	blogMaps := make(map[string][]*Blog)

	// now := time.Now()
	// year := now.Year()
	// yearStr := strconv.Itoa(year)
	// month := int(now.Month())
	// monthStr := strconv.Itoa(month)
	// firstDateOfMonth := yearStr + "-" + monthStr + "-01"

	var blogs []*Blog
	var blogKeys []string
	// qs.Filter("StatusId", 1).Filter("CreatedAt__gte", firstDateOfMonth).OrderBy("-CreatedAt").All(&blogs)
	qs.Filter("StatusId", 1).OrderBy("-CreatedAt").All(&blogs)

	if len(blogs) > 0 {
		currentKey := generateKey(blogs[0].CreatedAt)
		blogKeys = append(blogKeys, currentKey)
		tempBlogs := make([]*Blog, 0)
		tempBlogs = append(tempBlogs, blogs[0])

		blogMaps[currentKey] = tempBlogs
		for i, blog := range blogs {
			if i >= 1 {
				tempKey := generateKey(blog.CreatedAt)
				if currentKey == tempKey {
					tempBlogs = append(tempBlogs, blog)
					blogMaps[currentKey] = tempBlogs
				} else {
					currentKey = tempKey
					blogKeys = append(blogKeys, currentKey)
					tempBlogs = make([]*Blog, 0)
					tempBlogs = append(tempBlogs, blog)
					blogMaps[currentKey] = tempBlogs
				}
			}
		}
	}
	return blogKeys, blogMaps
}

func generateKey(t time.Time) string {

	key := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	return key
}
